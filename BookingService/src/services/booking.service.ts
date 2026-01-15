import PrismaClient from "../prisma/client";
import { CreateBookingDTO } from "../dto/booking.dto";
import { confirmBooking, createBooking, createIdempotencyKey, finalizeIdempotencyKey, getIdempotencyKeyWithLock } from "../repositories/booking.repository";
import { BadRequestError, InternalServerError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";
import { serverConfig } from "../config";
import { redlock } from "../config/redis.config";

export async function createBookingService(createBookingDTO: CreateBookingDTO) {

    const ttl = serverConfig.LOCK_TTL;
    const bookingResource = `hotel:${createBookingDTO.hotelId}`;

    try {
        await redlock.acquire([bookingResource], ttl);
        const booking = await createBooking({
            userId : createBookingDTO.userId,
            hotelId : createBookingDTO.hotelId,
            bookingAmount : createBookingDTO.bookingAmount,
            totalGuests : createBookingDTO.totalGuests
        })

        const idempotencyKey = generateIdempotencyKey();

        await createIdempotencyKey(idempotencyKey, booking.id);

        return {
            bookingId: booking.id,
            idempotencyKey: idempotencyKey
        }
    } catch (error) {
        throw new InternalServerError('Failed to acquire lock for booking resource');
    }
    
}

export async function confirmBookingService(idempotencyKey: string) {

    return await PrismaClient.$transaction(async(tx) => {
        const idempotencyKeyData = await getIdempotencyKeyWithLock(tx, idempotencyKey);

        if(!idempotencyKeyData || !idempotencyKeyData.bookingId) {
            throw new NotFoundError('idempotency key not found')
        }
        if(idempotencyKeyData.finalized) {
            throw new BadRequestError('idempotency key already finalized');
        }

        const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);
        await finalizeIdempotencyKey(tx, idempotencyKey);

        return booking;
    })

    
}
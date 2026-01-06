import { Request, Response, NextFunction } from "express";
import { confirmBookingService, createBookingService } from "../services/booking.service";
import { StatusCodes } from "http-status-codes";


export async function createBookingHandler(req: Request, res: Response, next: NextFunction) {
    const booking = await createBookingService(req.body);

    res.status(StatusCodes.CREATED).send({
        bookingId : booking.bookingId,
        idempotencyKey : booking.idempotencyKey
    })
}

export async function confirmBookingHandler(req: Request, res: Response, next: NextFunction) {
    const booking = await confirmBookingService(req.params.idempotencyKey);

    res.status(StatusCodes.OK).send({
        bookingId : booking.id,
        status : booking.status
    })
}
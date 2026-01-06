import { z } from "zod";

export const createBookingSchema = z.object({
    userId : z.number({message : "User ID must be present"}),
    hotelId : z.number({message : "Hotel ID must be present"}),
    bookingAmount : z.number({message : "booking amount must be present"}).min(1, {message : "booking amount cannot be 0"}),
    totalGuests : z.number({message : "totalGuests must be present"}).min(1, {message : "total guests cannot be 0"})
})
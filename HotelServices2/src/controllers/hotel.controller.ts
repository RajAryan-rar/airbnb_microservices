import { NextFunction, Request, Response } from "express";
import {  createHotelService, deleteHotelService, getAllHotelsService, getHotelByIdService } from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";

export async function createHotelHandler(req : Request, res: Response, Next : NextFunction) {
    //1. Call the service layer (delegation of logic)
    const hotelResponse = await createHotelService(req.body);

    //2. Send the response (response formulation)
    res.status(StatusCodes.CREATED).json({
        message : "Hotel created successfully",
        data : hotelResponse,
        success : true
    })
}

export async function getHotelByIdHandler(req: Request, res:Response, Next : NextFunction) {
    //1. handle the payload to the appropriate method in the service layer 
    const hotelResponse = await getHotelByIdService(Number(req.params.id));

    //2. Response formulation
    res.status(StatusCodes.OK).json({
        message : "hotel found",
        data : hotelResponse,
        success : true
    })
}

export async function getAllHotelsHandler(req: Request, res: Response, Next: NextFunction) {
    // call the service layer
    const hotelResponse = await getAllHotelsService();

    //send the response
    res.status(StatusCodes.OK).json({
        message : "Hotels found successfully",
        data : hotelResponse,
        success : true
    })
}


export async function deleteHotelHandler(req: Request, res: Response, Next: NextFunction) {
    const hotelResponse = await deleteHotelService(Number(req.params.id));

    res.status(StatusCodes.OK).json({
        message : "Hotel deleted successfully",
        data : hotelResponse,
        success : true
    })
}


// export async function updateHotelByIdController(req: Request, res: Response, Next: NextFunction) {
//     const hotelResponse = await updateHotelByIdService(Number(req.params.id), req.body.hotelName);
    
//     res.status(StatusCodes.OK).json({
//         message : "Hotel updated successfully",
//         data : hotelResponse,
//         success : true
//     })
// }
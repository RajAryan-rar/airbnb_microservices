import { createHotelDto } from "../dto/hotel.dto";
import { createHotel, getAllHotels, getHotelById, softDeleteHotel, updateHotelById } from "../repositories/hotel.repository";

export async function createHotelService(hotelData : createHotelDto) {
    const hotel = await createHotel(hotelData);
    return hotel;
}

export async function getHotelByIdService(id : number) {
    const hotel = await getHotelById(id);
    return hotel;
}

export async function getAllHotelsService() {
    const hotels = await getAllHotels();
    return hotels;
}

export async function deleteHotelService(id : number) {
    const response = await softDeleteHotel(id);
    return response;
}

export async function updateHotelByIdService(id: number, hotelName: string) {
    const response = await updateHotelById(id, hotelName);
    return response;
}
import { createHotelDto } from "../dto/hotel.dto";
// import { createHotel, getAllHotels, getHotelById, softDeleteHotel, updateHotelById } from "../repositories/hotel.repository";
import { HotelRepository } from "../repositories/hotel.repository";

const hotelRepository = new HotelRepository();

export async function createHotelService(hotelData : createHotelDto) {
    const hotel = await hotelRepository.create(hotelData);
    return hotel;
}

export async function getHotelByIdService(id : number) {
    const hotel = await hotelRepository.findById(id);
    return hotel;
}

export async function getAllHotelsService() {
    const hotels = await hotelRepository.findAll();
    return hotels;
}

export async function deleteHotelService(id : number) {
    const response = await hotelRepository.softDelete(id);
    return response;
}


import { createRoomCategoryDto } from "../dto/roomCategory.dto";
import { HotelRepository } from "../repositories/hotel.repository";
import { RoomCatRepository } from "../repositories/roomCategory.repository";
import { NotFoundError } from "../utils/errors/app.error";

const roomCategoryRepository = new RoomCatRepository();
const hotelRepository = new HotelRepository();

export async function createRoomCategoryService(roomCategoryData: createRoomCategoryDto) {
    const roomCategory = await roomCategoryRepository.create(roomCategoryData);
    return roomCategory;
}

export async function getRoomCategoryByIdService(id : number) {
    const roomCategory = await roomCategoryRepository.findById(id);
    return roomCategory;
}

export async function getAllRoomCategoriesByHotelIdService(id : number) {
    const hotel = await hotelRepository.findById(id);

    if(!hotel) {
        throw new NotFoundError(`Hotel with id : ${id} not found`);
    }

    const roomCategories = await roomCategoryRepository.findAllByHotelId(id);
    return roomCategories;
}

export async function deleteRoomCategoryService(id : number) {
    const roomCategory = await roomCategoryRepository.softDelete(id);
    return roomCategory;
}

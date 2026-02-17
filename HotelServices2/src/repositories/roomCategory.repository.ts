import logger from "../config/logger.config";
import RoomCategory from "../db/models/roomCategory";
import { NotFoundError } from "../utils/errors/app.error";
import BaseRepository from "./base.repository";

export class RoomCatRepository extends BaseRepository<RoomCategory> {
    constructor() {
        super(RoomCategory);
    }

    async findAllByHotelId(id : number) {
        const roomCategories = await this.model.findAll({
            where : {
                hotelId : id,
                deletedAt : null
            }
        });

        if(!roomCategories || roomCategories.length === 0) {
            throw new NotFoundError(`No roomCategory found with hotelId : ${id}`);
        }

        return roomCategories;
    }

    async softDelete(id: number) {
        const roomCategory = await this.model.findByPk(id);

        if(!roomCategory) {
            throw new NotFoundError(`room category with id : ${id} not found`);
        }

        roomCategory.deletedAt = new Date();
        await roomCategory.save();
        logger.info(`room category soft deleted : ${id}`);
        return roomCategory;
    }
}
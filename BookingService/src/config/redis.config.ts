import IOredis, {Redis} from 'ioredis';
import Redlock from 'redlock';
import { serverConfig } from '.';

export const redisClient = new IOredis(serverConfig.REDIS_SERVER_URL);


function connectToRedis() {
    try {
        let connection: Redis;

       
        return () => {
            if(!connection) {
                connection = new IOredis(serverConfig.REDIS_SERVER_URL);
                return connection;
            }
            return connection;
        }

    } catch (error) {
        console.log("Error connecting to Redis", error);
        throw error;
    }
}

export const getRedisConnObject = connectToRedis();

export const redlock = new Redlock([getRedisConnObject()], {
    driftFactor: 0.01,
    retryCount: 10,
    retryDelay: 200,
    retryJitter: 200,
    // automaticExtensionThreshold: 500,
}) 
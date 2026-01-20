import { Job, Worker } from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producer";


export const setupMailerWorker = () => {
    const emailProcessor = new Worker<NotificationDto>(
    MAILER_QUEUE, //name of my queue
    async (job: Job) => {
        if(job.name !== MAILER_PAYLOAD) {
            throw new Error('Invalid job object')
        }
    }, //process fxn
        {
            connection: getRedisConnObject(),
        }
    )

    emailProcessor.on('failed', () => {
        console.error('Email processing failed');
    })

    emailProcessor.on('completed', () => {
        console.log('Email processing completed successfully');
    })
}
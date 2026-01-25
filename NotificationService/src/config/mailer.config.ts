import nodemailer from 'nodemailer';
import { serverConfig } from '.';

const transporter = nodemailer.createTransport({
    service: 'gmail',
    auth : {
        pass: serverConfig.MAIL_PASS,
        user: serverConfig.MAIL_USER
    }
});

export default transporter;
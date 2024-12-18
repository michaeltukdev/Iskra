import { ErrorObject } from "@/app/types/utils";

export const splitError = (error: string): ErrorObject => {
    const errorObject: ErrorObject = {};

    const errorArray = error.split(';').filter(err => err.trim() !== '');

    errorArray.forEach((err) => {
        const [key, value] = err.split(':').map(part => part?.trim());
        if (key && value) {
            errorObject[key] = value;
        } else if (key) {
            errorObject.error = key;
        }
    });

    if (Object.keys(errorObject).length === 0) {
        errorObject.error = error;
    }

    return errorObject;
};

export function sleep(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}
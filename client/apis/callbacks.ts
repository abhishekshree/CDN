import axios from 'axios';

const BASE_URL = 'http://localhost:8080/api';
const UPLOAD = BASE_URL + '/upload';
const VIEW = BASE_URL + '/view';
const VIEWALL = BASE_URL + '/view/all';
const DELETE = BASE_URL+ '/delete';
const ZIP = BASE_URL+ '/zip';


export const postUpload = async (f: FormData) => {
    try {
        const response = await axios.post(UPLOAD, f);
        console.log(response);
        return response.data;
    }
    catch (error) {
        console.log(error);
        return error;
    }
}

export const getViewAll = async () => {
    try {
        const response = await axios.get(VIEWALL);
        console.log(response);
        return response.data;
    }
    catch (error) {
        console.log(error);
        return error;
    }
}

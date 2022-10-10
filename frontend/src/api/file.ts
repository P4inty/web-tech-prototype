import axios from 'axios';
import type { MetaData, MetaDataResponse } from '../types/file.type';

export class FileApi {

    static _BASE_URL = import.meta.env.VITE_API_URL

    static async uploadMetadata(data: MetaData): Promise<MetaDataResponse> {
        return await axios.post(this._BASE_URL + "/upload/meta", data)
    }

    static async uploadFile(file: File, uri: String) {
        const data = new FormData()
        data.append("file", file)
        return await axios.post(this._BASE_URL + `/upload/${uri}`, data, {
            headers: {
                'Content-Type': 'multipart/form-data',
            }
        })
    }
}
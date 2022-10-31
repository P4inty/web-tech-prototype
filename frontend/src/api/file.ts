import axios from 'axios';
import type { CustomFilesResponse, MetaData, MetaDataResponse } from '../types/file.type';

export class FileApi {

    static BASE_URL = import.meta.env.VITE_API_URL?.length > 0 ? import.meta.env.VITE_API_URL : "http://" + window.location.hostname + import.meta.env.VITE_API_PORT

    static async uploadMetadata(data: MetaData): Promise<MetaDataResponse> {
        console.log(this.BASE_URL + "/upload/meta")
        return await axios.post(this.BASE_URL + "/upload/meta", data)
    }

    static async uploadFile(file: File, uri: String): Promise<any> {
        const data = new FormData()
        data.append("file", file)
        return await axios.post(this.BASE_URL + `/upload/${uri}`, data, {
            headers: {
                'Content-Type': 'multipart/form-data',
            }
        })
    }

    static async getFiles(): Promise<CustomFilesResponse> {
        return await axios.get(this.BASE_URL + "/uploads")
    }
}
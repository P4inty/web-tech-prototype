import type { Tag } from './tag.type'


export interface MetaData {
    Name: string
    Description: string
    Tags: Tag[]
}

export interface MetaDataResponse {
    data: {
        uri: string
    }
}

export interface CustomFile extends MetaData {
    ID: number
    Uri: string
    CreatedAt: string
}

export interface CustomFilesResponse {
    data: {
        files: CustomFile[]
    }
}

export interface DownloadFileResponse {
    data: Blob
}
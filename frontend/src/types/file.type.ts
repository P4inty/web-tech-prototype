import type { Tag } from './tag.type'


export interface MetaData {
    Name: String
    Description: String
    Tags: Tag[]
}

export interface MetaDataResponse {
    data: {
        uri: String
    }
}

export interface CustomFile extends MetaData {
    ID: Number
    Uri: String
    CreatedAt: Date
}

export interface CustomFilesResponse {
    data: {
        files: CustomFile[]
    }
}

export interface DownloadFileResponse {
    data: Blob
}
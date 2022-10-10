import type { Tag } from './tag.type'


export interface MetaData {
    name: String,
    description: String,
    tags: Tag[]
}

export interface MetaDataResponse {
    data: {
        uri: String
    }
}
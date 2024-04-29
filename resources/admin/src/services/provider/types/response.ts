export interface Response<Data extends object = object> {
    code: number
    msg: string
    data: Data
    meta: Meta
}

export interface Meta {
    pagination: Pagination
}

export interface Pagination {
    current_page: number
    per_page: number
    total: number
}

export type ListResponse<RecordType extends object = object> = Response<RecordType[]>
export type ShowResponse<RecordType extends object = object> = Response<RecordType>

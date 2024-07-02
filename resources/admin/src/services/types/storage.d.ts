interface Storage {
    id: number
    name: string
    driver: string
    thumbnail: number
    option: Option
    created_at: string
    updated_at: string
}

interface Option {
    local?: Local
}

interface Local {
    folder: string
}

interface FileDescription {
    id: number
    is_dir: boolean
    path: string
    filename: string
    mode: string
    size: number
    created_at: string
    accessed_at: string
    updated_at: string

    exif?: ExifItem[]
}

interface ExifItem {
    key: string
    label: string
    value: string
    enums: Enums
}

type Enums = Record<string, string>

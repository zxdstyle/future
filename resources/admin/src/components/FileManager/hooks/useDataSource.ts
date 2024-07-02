import { useInfiniteList } from "@refinedev/core"
import { useContext } from "react"
import { StorageContext } from "../Content/context"

export default function useDataSource() {
    const { currentPath } = useContext(StorageContext)
    const { data } = useInfiniteList<FileDescription>({
        resource: "fs",
        filters: [
            { field: "parent_dir", operator: "eq", value: currentPath },
        ],
        queryOptions: {
            enabled: !!currentPath,
        },
    })

    const pages = data ? data.pages.map(page => page.data) : []
    const allPages: FileDescription[] = data ? [].concat(...pages) : []

    return { data: allPages }
}

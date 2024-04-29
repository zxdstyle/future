import { useSearchParams } from "react-router-dom"
import { useInfiniteList } from "@refinedev/core"

export default function useDataSource() {
    const [params] = useSearchParams()

    const { data } = useInfiniteList<FileDescription>({
        resource: "storage",
        filters: [
            { field: "parent_dir", operator: "eq", value: params.get("parent_dir") },
        ],
    })

    const pages = data ? data.pages.map(page => page.data) : []
    // @ts-expect-error
    const allPages: FileDescription[] = data ? [].concat(...pages) : []

    return { data: allPages }
}

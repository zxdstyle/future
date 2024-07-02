import type { BreadcrumbProps } from "antd"
import { Breadcrumb } from "antd"
import { Link, useSearchParams } from "react-router-dom"
import Folder from "@/assets/icons/folder.png"

export default function Footer() {
    const [params] = useSearchParams()
    const items = params.get("parent_dir")?.toString().split("/").filter(item => !!item).map((item) => {
        return {
            path: item,
            title: <div className="flex items-center gap-0.5 text-gray-400">
                <img src={Folder} alt={item} className="w-4" />
                {item}
            </div>,
        }
    })

    const itemRender: BreadcrumbProps["itemRender"] = (current, _params, _items, paths) => {
        return <Link to={`/storage?parent_dir=/${paths.join("/")}`}>{current.title}</Link>
    }

    return (
        <footer className="h-12 w-full border-t border-gray-800 border-solid flex items-center justify-between px-5">
            <Breadcrumb items={items} itemRender={itemRender} />
        </footer>
    )
}

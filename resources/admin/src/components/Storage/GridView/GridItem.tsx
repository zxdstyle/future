import { useNavigate } from "react-router-dom"
import type { MouseEvent } from "react"
import type { BaseRecord } from "@refinedev/core"
import ContextMenu from "@/components/Storage/ContextMenu"
import FileThumbnail from "@/components/Storage/FileThumbnail"
import { humanReadableFilesize } from "@/utils/human.ts"

interface IGridItemProps {
    item: FileDescription
}

export default function GridItem({ item }: IGridItemProps) {
    const navigate = useNavigate()
    const handleDbClick = (e: MouseEvent<HTMLLIElement>, file: BaseRecord) => {
        e.stopPropagation()
        e.preventDefault()

        if (file.is_dir)
            navigate(`/storage?parent_dir=${file.path}`)
    }

    return (
        <ContextMenu item={item}>
            <li
                key={item.filename}
                className="w-24 h-36 text-center"
                onDoubleClick={e => handleDbClick(e, item)}
            >

                <div className="rounded-xl p-2 hover:bg-slate-700">
                    <FileThumbnail width={80} item={item} />
                </div>

                <div className="text-sm break-words whitespace-normal line-clamp-2 text-white">{item.filename}</div>
                <div className="text-[10px] text-primary/60">{humanReadableFilesize(item.size)}</div>

            </li>
        </ContextMenu>
    )
}

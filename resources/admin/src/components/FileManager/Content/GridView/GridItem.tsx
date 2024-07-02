import type { MouseEvent } from "react"
import { useContext } from "react"
import type { BaseRecord } from "@refinedev/core"
import ContextMenu from "../ContextMenu"
import FileThumbnail from "../FileThumbnail"
import { StorageContext } from "../context"
import { humanReadableFilesize } from "@/utils/human"

interface IGridItemProps {
    item: FileDescription
}

export default function GridItem({ item }: IGridItemProps) {
    const { setCurrentPath } = useContext(StorageContext)
    const handleDbClick = (e: MouseEvent<HTMLLIElement>, file: BaseRecord) => {
        e.stopPropagation()
        e.preventDefault()

        if (file.is_dir)
            setCurrentPath(file.path)
    }

    return (
        <ContextMenu item={item}>
            <li
                key={item.filename}
                className="w-20 text-center flex flex-col items-center"
                onDoubleClick={e => handleDbClick(e, item)}
            >

                <div className="rounded-xl hover:bg-slate-700">
                    <FileThumbnail width={80} item={item} />
                </div>

                <div className="text-sm break-words whitespace-normal line-clamp-2 text-white">{item.filename}</div>
                <div className="text-[10px] text-primary/60">{humanReadableFilesize(item.size)}</div>

            </li>
        </ContextMenu>
    )
}

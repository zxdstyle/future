import clsx from "clsx"
import { useContext } from "react"
import { NavbarContext } from "./context"
import { tools } from "./components"

interface Props {
    className?: string
}

export default function Navbar({ className }: Props) {
    const { leftTools } = useContext(NavbarContext)

    return (
        <nav className={clsx("h-14 px-5 w-full border-b border-gray-800 border-solid flex items-center justify-between", className)}>
            <div className="flex items-center gap-6">
                {leftTools.map((item) => {
                    return <span key={item}>{tools[item]}</span>
                })}
            </div>
        </nav>
    )
}

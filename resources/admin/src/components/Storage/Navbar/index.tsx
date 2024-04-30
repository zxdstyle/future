import { Button } from "antd"
import clsx from "clsx"
import { useNavigate } from "react-router-dom"
import QuickJumper from "./QuickJumper"

interface Props {
    className?: string
}

export default function Navbar({ className }: Props) {
    const navigate = useNavigate()
    return (
        <nav className={clsx("h-14 px-5 w-full border-b border-gray-800 border-solid flex items-center justify-between", className)}>
            <div className="flex items-center gap-6">
                <Button.Group>
                    <Button onClick={() => navigate(-1)} type="text" icon={<IconTablerArrowNarrowLeft className="text-2xl" />}></Button>
                    <Button onClick={() => navigate(+1)} type="text" icon={<IconTablerArrowNarrowLeft className="rotate-180 text-2xl" />}></Button>
                </Button.Group>

                <div>
                    <QuickJumper />
                </div>
            </div>
        </nav>
    )
}

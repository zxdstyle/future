import { useContext, useEffect } from "react"
import StorageManager from "@/components/Storage"
import { NavbarContext } from "@/components/layout/Navbar/context.tsx"

export function Component() {
    const { setLeftTools } = useContext(NavbarContext)
    useEffect(() => {
        setLeftTools(["navigator", "quick-jumper"])
    }, [])

    return (
        <div className="h-full">
            <StorageManager />
        </div>
    )
}

import { createContext, useState } from "react"
import type { ToolKey } from "./components"

interface Props {
    leftTools: string[]
    setLeftTools: (value: ToolKey[]) => void
    rightTools: string[]
    setRightTools: (value: ToolKey[]) => void
}

export const NavbarContext = createContext<Props>({
    leftTools: [],
    rightTools: [],
    setLeftTools: () => {},
    setRightTools: () => {},
})

export function NavbarProvider({ children }) {
    const [leftTools, setLeftTools] = useState([])
    const [rightTools, setRightTools] = useState([])

    return (
        <NavbarContext.Provider
            value={{
                leftTools,
                setLeftTools,
                rightTools,
                setRightTools,
            }}
        >
            {children}
        </NavbarContext.Provider>
    )
}

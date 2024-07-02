import type { PropsWithChildren } from "react"
import { createContext, useEffect, useState } from "react"
import { ConfigProvider } from "antd"
import light from "@/config/theme/light"
import dark from "@/config/theme/dark"

type Mode = "dark" | "light" | "system"

interface ColorModeContextType {
    mode: "dark" | "light"
    setMode: (mode: "dark" | "light") => void
}

export const ColorModeContext = createContext<ColorModeContextType>({} as ColorModeContextType)

export function ColorModeContextProvider({ children }: PropsWithChildren) {
    const colorModeFromLocalStorage = localStorage.getItem("colorMode") as Mode

    const isSystemPreferenceDark = window?.matchMedia("(prefers-color-scheme: dark)").matches

    const systemPreference = isSystemPreferenceDark ? "dark" : "light"

    const theme = colorModeFromLocalStorage === "system" ? systemPreference : colorModeFromLocalStorage

    const [mode, setMode] = useState<"dark" | "light">(theme || "dark")

    useEffect(() => {
        window.localStorage.setItem("colorMode", mode)
    }, [mode])

    const setColorMode = (mode: "dark" | "light") => {
        setMode(mode)
    }

    return (
        <ColorModeContext.Provider
            value={{
                setMode: setColorMode,
                mode,
            }}
        >
            <ConfigProvider
                theme={mode === "light" ? light : dark}
            >
                {children}
            </ConfigProvider>
        </ColorModeContext.Provider>
    )
}

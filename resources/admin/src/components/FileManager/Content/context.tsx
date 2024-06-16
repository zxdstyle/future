import type { PropsWithChildren } from "react"
import { createContext, useEffect, useState } from "react"
import type { OverviewProps } from "../hooks/useOverview"
import { useOverview } from "../hooks/useOverview"

export interface IStorageContext extends OverviewProps {
    rootFolder: string
    currentPath: string
    setCurrentPath: (path: string) => void
}

export const StorageContext = createContext<IStorageContext>({} as IStorageContext)

interface IStorageProviderProps {
    root: string
}
export function Provider({ children, root }: PropsWithChildren<IStorageProviderProps>) {
    const { openOverview, closeOverview, showOverview, description } = useOverview()

    // const [rootFolder, setRootFolder] = useState(root)
    const [currentPath, setCurrentPath] = useState("")
    useEffect(() => setCurrentPath(root), [root])

    return (
        <StorageContext.Provider
            value={{
                openOverview,
                closeOverview,
                showOverview,
                description,
                rootFolder: root,
                currentPath,
                setCurrentPath,
            }}
        >
            {children}
        </StorageContext.Provider>
    )
}

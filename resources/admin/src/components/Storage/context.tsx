import type { PropsWithChildren } from "react"
import { createContext } from "react"
import type { OverviewProps } from "./hooks/useOverview"
import { useOverview } from "./hooks/useOverview"

export interface IStorageContext extends OverviewProps {}

export const StorageContext = createContext<IStorageContext>({
    openOverview: () => {},
    closeOverview: () => {},
    showOverview: false,
})

export function Provider({ children }: PropsWithChildren) {
    const { openOverview, closeOverview, showOverview, description } = useOverview()

    return <StorageContext.Provider value={{ openOverview, closeOverview, showOverview, description }}>{children}</StorageContext.Provider>
}

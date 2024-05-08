import type { RouteObject } from "react-router-dom"

const albumRoutes: RouteObject[] = [
    { path: "/album", lazy: () => import("@/views/album") },
]

export default albumRoutes

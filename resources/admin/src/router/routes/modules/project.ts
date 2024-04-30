import type { RouteObject } from "react-router-dom"

const projectRoutes: RouteObject[] = [
    { path: "/project", lazy: () => import("@/views/project") },
]

export default projectRoutes

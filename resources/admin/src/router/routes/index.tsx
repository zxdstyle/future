import type { RouteObject } from "react-router-dom"
import { ErrorComponent } from "@refinedev/antd"
import { handleModuleRoutes } from "../utils"
import { Provider, ThemedLayoutV2 } from "@/components/layout"

const modules = import.meta.glob("./modules/**/*.ts", { eager: true }) as Record<string, { default: RouteObject }>

const routes: RouteObject[] = [
    {
        path: "/",
        element: <Provider />,
        children: [
            {
                path: "/",
                element: <ThemedLayoutV2 />,
                children: handleModuleRoutes(modules),
            },
            { path: "*", element: <ErrorComponent /> },
        ],
    },
]

export default routes

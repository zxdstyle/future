import { RouterProvider as InnerRouterProvider, createBrowserRouter } from "react-router-dom"
import routes from "./routes"

const router = createBrowserRouter(routes)

export default function RouterProvider() {
    return <InnerRouterProvider router={router} />
}

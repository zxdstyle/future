import "@refinedev/antd/dist/reset.css"
import "./assets/styles/style.less"

import { ConfigProvider } from "antd"
import zhCN from "antd/locale/zh_CN"
import dayjs from "dayjs"
import RouterProvider from "@/router"

import "dayjs/locale/zh-cn"

dayjs.locale("zh-cn")

export default function App() {
    return (
        <ConfigProvider locale={zhCN}>
            <RouterProvider />
        </ConfigProvider>
    )
}

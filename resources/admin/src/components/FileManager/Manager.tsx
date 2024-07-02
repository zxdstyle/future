import type { ThemeConfig } from "antd"
import { ConfigProvider, Flex, Modal, theme } from "antd"
import { useState } from "react"
import type { FileManagerProps } from "./types/manager"
import Sidebar from "./Sidebar"
import Content from "./Content"

const modalStyle: ThemeConfig = {
    components: {
        Modal: {
            paddingMD: 0,
            paddingContentHorizontalLG: 0,
        },
    },
}

export default function Manager({ modalProps }: FileManagerProps) {
    const { token } = theme.useToken()
    const [root, setRoot] = useState("")

    return (
        <ConfigProvider theme={modalStyle}>
            <Modal {...modalProps}>
                <Flex className="h-128">
                    <Sidebar style={{ borderRight: `1px solid ${token.colorSplit}` }} onChange={s => setRoot(s.option.local.folder)} />

                    <Content root={root} />
                </Flex>
            </Modal>
        </ConfigProvider>
    )
}

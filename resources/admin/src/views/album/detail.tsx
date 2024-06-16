import { useContext, useEffect } from "react"
import { Button, Flex } from "antd"
import { NavbarContext } from "@/components/layout"
import { FileManager } from "@/components/FileManager"
import { useFileManager } from "@/components/FileManager/hooks/useFileManager.tsx"

export function Component() {
    const { setLeftTools } = useContext(NavbarContext)
    useEffect(() => {
        setLeftTools(["navigator"])
    }, [])

    const { fileManagerProps, show } = useFileManager({ defaultOpen: true })

    return (
        <>
            <Flex align="center" justify="center" className="w-full h-full">
                <Button className="w-52 h-52 flex flex-col" onClick={show} icon={<IconTablerUpload className="text-5xl" />}>
                    选择照片
                </Button>
            </Flex>

            <FileManager {...fileManagerProps} />
        </>

    )
}

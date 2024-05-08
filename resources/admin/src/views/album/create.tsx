import { useContext, useEffect } from "react"
import type { SaveButtonProps } from "@refinedev/antd"
import { useDrawerForm } from "@refinedev/antd"
import type { DrawerProps, FormProps } from "antd"
import { Button, Drawer } from "antd"
import { NavbarContext } from "@/components/layout"
import Form from "@/views/album/form.tsx"

export default function Component() {
    const { setLeftTools } = useContext(NavbarContext)
    useEffect(() => {
        setLeftTools(["navigator"])
    }, [])

    const { drawerProps, show, formProps, saveButtonProps } = useDrawerForm({
        resource: "albums",
        action: "create",
    })

    return (
        <>
            <Button
                onClick={() => show()}
                className="h-40 !w-32 flex flex-col items-center justify-center"
                icon={<IconTablerPlus className="text-3xl" />}
            >
                添加相册
            </Button>

            <DrawerContent formProps={formProps} drawerProps={drawerProps} saveButtonProps={saveButtonProps} />
        </>
    )
}
interface Props {
    formProps: FormProps
    drawerProps: DrawerProps
    saveButtonProps: SaveButtonProps

}
function DrawerContent({ formProps, drawerProps, saveButtonProps }: Props) {
    return (
        <Drawer {...drawerProps}>
            <Form {...formProps} saveButtonProps={saveButtonProps} />
        </Drawer>
    )
}

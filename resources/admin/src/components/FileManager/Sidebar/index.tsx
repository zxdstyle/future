import clsx from "clsx"
import type { CSSProperties } from "react"
import type { MenuProps, ThemeConfig } from "antd"
import { Button, ConfigProvider, Flex, Image, Menu } from "antd"
import { useList } from "@refinedev/core"
import CreateStorage from "../Storage/create"
import StorageImg from "@/assets/icons/storage.png"
import HDD from "@/assets/icons/hdd.png"

const menuStyle: ThemeConfig = {
    token: {
    },
    components: {
        Menu: {
            itemHeight: 25,
            padding: 0,
            paddingXS: 0,
            fontSize: 12,
            borderRadius: 0,
            activeBarBorderWidth: 0,
        },
    },
}

interface Props {
    className?: string
    style?: CSSProperties
    onChange?: (storage: Storage) => void
}

export default function ({ className, style, onChange }: Props) {
    const { data } = useList<Storage>({
        resource: "storages",
    })

    const items: MenuProps["items"] = [
        {
            type: "group",
            label: <Flex align="center" justify="space-between" gap={4}>
                <span>
                    <Image src={StorageImg} width={20} preview={false} />
                    <span>存储</span>
                </span>
                <Button icon={<IconTablerDots />} size="small" type="text" href="/settings/storage" />
            </Flex>,
            children: data?.data.map((datum) => {
                return {
                    label: <Flex gap={4}>
                        <Image src={HDD} width={20} height={20} preview={false} />
                        <span>{datum.name}</span>
                    </Flex>,
                    key: datum.id,
                }
            }),
        },
    ]

    const handleSelect: MenuProps["onSelect"] = (info) => {
        if (data) {
            const storages = data.data.filter(datum => String(datum.id) === info.key)
            if (storages && storages.length > 0)
                onChange && onChange(storages[0])
        }
    }

    return (
        <div
            style={style}
            className={clsx(className, "relative flex w-44 flex-shrink-0 flex-col gap-2.5 p-3")}
        >
            <div className="no-scrollbar flex flex-col space-y-5 overflow-x-hidden overflow-y-scroll">
                <ConfigProvider theme={menuStyle}>
                    <Menu items={items} selectable onSelect={handleSelect} />
                </ConfigProvider>

                <CreateStorage>
                    <Button block ghost className="text-xs border-dashed" size="small">添加存储</Button>
                </CreateStorage>
            </div>
        </div>
    )
}

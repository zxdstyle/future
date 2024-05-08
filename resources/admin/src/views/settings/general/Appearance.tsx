import { Flex, InputNumber, Slider } from "antd"

export default function Appearance() {
    return (
        <Flex justify="space-between" className="w-full text-base mt-4">
            <div>
                <h3 className="text-white text-base">CPU</h3>
                <small>后台处理时可以使用CPU的量</small>
            </div>
            <div className="flex items-end gap-2">
                <Slider min={1} max={100} step={10} className="w-44" tooltip={{ formatter: val => `${val}%` }} />
                <InputNumber min={1} max={100} controls={false} width={10} />
            </div>
        </Flex>
    )
}

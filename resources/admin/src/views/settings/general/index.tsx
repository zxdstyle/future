import CPU from "./Cpu"
import Appearance from "./Appearance"

export function Component() {
    return (
        <div className="px-12 w-full">
            <h1 className="text-white text-2xl mb-2">通用设置</h1>
            <h2>一些通用的一般设置</h2>

            <CPU />

            <Appearance />
        </div>
    )
}

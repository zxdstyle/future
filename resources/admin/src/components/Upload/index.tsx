import type { UploadFile, UploadProps } from "antd"
import { Button, Divider, Flex, Progress, Upload, theme } from "antd"
import { useState } from "react"

export default function UploadModal() {
    const [files, setFiles] = useState<UploadFile[]>([])

    const handleChange: UploadProps["onChange"] = (info) => {
        setFiles([...info.fileList])
        console.log(info.fileList)
    }

    const { token } = theme.useToken()

    return (
        <div className="min-w-160 min-h-96">
            <Flex justify="end">
                <Upload onChange={handleChange} fileList={files} showUploadList={false} accept="image/*" multiple>
                    <Button icon={<IconTablerUpload />}>上传文件</Button>
                </Upload>
            </Flex>

            <Divider className="my-3" />
            <ul>
                {files.filter(file => file.status !== "removed").map(file => (
                    <li key={file.uid} className={file.status === "error" ? "text-red-500" : ""}>
                        <span>{file.name}</span>
                        <Progress percent={file.percent} strokeColor={token.colorSuccess} />
                    </li>
                ))}
            </ul>
        </div>
    )
}

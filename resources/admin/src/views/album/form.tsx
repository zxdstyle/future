import type { FormProps } from "antd"
import { Form, Input } from "antd"
import { SaveButton, type SaveButtonProps } from "@refinedev/antd"

interface Props extends FormProps {
    saveButtonProps: SaveButtonProps
}
export default function AlbumForm({ saveButtonProps, ...rest }: Props) {
    return (
        <div>
            <Form layout="vertical" {...rest}>
                <Form.Item label="名称" name="title" rules={[{ required: true }]}>
                    <Input />
                </Form.Item>
                <Form.Item label="副标题" name="sub_title">
                    <Input />
                </Form.Item>
                <Form.Item label="描述" name="description">
                    <Input.TextArea rows={6} />
                </Form.Item>
                <Form.Item>
                    <SaveButton {...saveButtonProps}></SaveButton>
                </Form.Item>
            </Form>
        </div>
    )
}

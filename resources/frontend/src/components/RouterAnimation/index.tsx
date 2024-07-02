import { Motion, Presence, Options } from 'solid-motionone';
import { JSX, Show, splitProps, useContext } from 'solid-js';
import { RouterStateContext } from '@/layouts/context/router';

interface Props extends Options {
    children?: JSX.Element;
    class?: string;
    style?: JSX.CSSProperties;
}

export default function RouterAnimation(_props: Props) {
    const [props, rest] = splitProps(_props, ['children']);
    const { leaving } = useContext(RouterStateContext);
    return (
        <Presence exitBeforeEnter>
            <Show when={!leaving()}>
                <Motion.div {...rest}>{props.children}</Motion.div>
            </Show>
        </Presence>
    );
}

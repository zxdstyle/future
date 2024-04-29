import { Accessor, createContext, createSignal, ParentProps } from 'solid-js';
import { RouteSectionProps, useBeforeLeave } from '@solidjs/router';

export type RouterStateCtx = {
    leaving: Accessor<boolean>;
};

export const RouterStateContext = createContext<RouterStateCtx>({
    leaving: () => false,
});

export function RouterStateProvider(props: ParentProps) {
    const [leaving, setLeaving] = createSignal(false);
    useBeforeLeave(e => {
        e.preventDefault();
        setLeaving(true);
        console.log('attempting to leave');
        setTimeout(() => {
            setLeaving(false);
            e.retry(true);
            console.log('success to leaved');
        }, 700);
    });

    return <RouterStateContext.Provider value={{ leaving }}>{props.children}</RouterStateContext.Provider>;
}

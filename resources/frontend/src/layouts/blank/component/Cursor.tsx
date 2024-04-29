import { createSignal, JSX } from 'solid-js';
import { Motion } from 'solid-motionone';

import RouterAnimation from '@/components/RouterAnimation';
import Slider from '@/components/slider';
import { CursorProvider, CursorState } from '@/layouts/context/cursor';
import { createMousePosition } from '@solid-primitives/mouse';
import { spring } from 'motion';
import styles from '../styles/cursor.module.css';

export default function CursorToggleContext(props: { children: JSX.Element }) {
    const [cursorState, setCursorState] = createSignal<CursorState>('inactive');

    const setActive = () => {
        setCursorState('active');
    };

    const setInactive = () => {
        setCursorState('inactive');
    };

    const pos = createMousePosition();

    const images = [
        'https://images.pexels.com/photos/6954544/pexels-photo-6954544.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/6954509/pexels-photo-6954509.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/11153461/pexels-photo-11153461.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/5049271/pexels-photo-5049271.jpeg?auto=compress&cs=tinysrgb&w=1600',
    ];

    return (
        <CursorProvider value={{ cursorState, setActive, setInactive }}>
            <Motion.div
                initial={false}
                animate={{ x: pos.x - 16, y: pos.y - 16 }}
                transition={{ easing: spring() }}
                class={styles.cursor}
            >
                <div
                    classList={{
                        'w-8 h-8 absolute left-0 top-0 border-2 rounded-full transition-all': true,
                        'border-white/30 ': cursorState() === 'inactive',
                        'scale-200 border-white': cursorState() === 'active',
                    }}
                />
                <span />
            </Motion.div>
            <RouterAnimation
                class="absolute top-0 left-0 right-0 bottom-0 w-screen h-screen bg-black"
                exit={{ opacity: 0 }}
                style={{ background: 'radial-gradient(ellipse at left top, #28282e 0%, #000000 70%)' }}
                transition={{ duration: 1 }}
            >
                <Slider images={images} />
            </RouterAnimation>

            <div
                classList={{
                    'absolute top-0 left-0 w-full h-full pointer-events-none transition-all duration-700': true,
                    'bg-black/60': cursorState() === 'inactive',
                    'bg-black/80': cursorState() === 'active',
                }}
            />

            {props.children}
        </CursorProvider>
    );
}

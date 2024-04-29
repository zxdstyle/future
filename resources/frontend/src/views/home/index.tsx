import styles from './styles/index.module.css';
import clsx from 'clsx';
import Slider from '@/components/slider';
import NavLine from '@/components/NavLine';
import RouterAnimation from '@/components/RouterAnimation';

export default function () {
    return <NavLine title="浏览作品" href="/album" subTitle="Explore Works" position="center" class="left-1/3" />;
}

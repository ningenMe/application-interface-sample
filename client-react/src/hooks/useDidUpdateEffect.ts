import { useRef, useEffect, EffectCallback, DependencyList } from 'react';

// 初回の実行がスキップされるuseEffect
// https://zenn.dev/catnose99/scraps/30c623ba72d6b5
export function useDidUpdateEffect(fn: EffectCallback, deps: DependencyList) {
  const didMountRef = useRef(false);

  useEffect(() => {
    if (!didMountRef.current) {
      didMountRef.current = true;
    } else {
      fn();
    }
  }, deps);
}

declare module 'vue' {
    interface ComponentCustomProperties {
        $filePath: (typeof import('@/methods/index'))['filePath']
    }
}

export {}
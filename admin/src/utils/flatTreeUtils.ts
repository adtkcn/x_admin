export type WithIdPid = {
    id?: string | number
    pid?: string | number | null | undefined
}

// ---------- 1. 查询匹配节点 ----------
export function findMatchedNodes<T extends WithIdPid>(
    data: T[],
    keyword: string,
    fields: string[] = ['name'],
    exact = false
): T[] {
    if (!keyword) return []
    return data.filter((node) =>
        fields.some((field) => {
            const val = node[field]
            return typeof val === 'string' && (exact ? val === keyword : val.includes(keyword))
        })
    )
}

// ---------- 2. 查询所有子级（后代） ----------
export function findDescendants<T extends WithIdPid>(data: T[], rootIds: (string | number)[]): T[] {
    // 构建 pid -> children 映射
    const childrenMap = new Map<string | number, T[]>()
    for (const node of data) {
        const p = node.pid ?? '__null__'
        if (!childrenMap.has(p)) childrenMap.set(p, [])
        childrenMap.get(p)!.push(node)
    }

    const result: T[] = []
    const queue: (string | number)[] = [...rootIds]
    const visited = new Set<string | number>()

    while (queue.length > 0) {
        const id = queue.shift()!
        if (visited.has(id)) continue
        visited.add(id)

        const children = childrenMap.get(id) || []
        for (const child of children) {
            result.push(child)
            queue.push(child.id)
        }
    }

    return Array.from(new Map(result.map((n) => [n.id, n])).values())
}

// ---------- 3. 查询所有父级（祖先） ----------
export function findAncestors<T extends WithIdPid>(data: T[], startIds: (string | number)[]): T[] {
    // 构建 id -> node 映射
    const nodeMap = new Map<string | number, T>()
    for (const node of data) {
        nodeMap.set(node.id, node)
    }

    const result: T[] = []
    const visited = new Set<string | number>()

    for (const id of startIds) {
        let currentId: string | number | null | undefined = id
        while (currentId != null) {
            const node = nodeMap.get(currentId)
            if (!node || visited.has(currentId)) break

            visited.add(currentId)
            // 跳过起始节点自身（只查“父级”）
            if (currentId !== id) {
                result.push(node)
            }

            currentId = node.pid
        }
    }

    return Array.from(new Map(result.map((n) => [n.id, n])).values())
}

// ---------- 4. 合并所有层级 ----------
export function collectAllLevels<T extends WithIdPid>(
    matched: T[],
    ancestors: T[],
    descendants: T[]
): T[] {
    const all = [...matched, ...ancestors, ...descendants]
    return Array.from(new Map(all.map((n) => [n.id, n])).values())
}

// ---------- 5. 主入口：一键获取结构化结果 ----------
export function queryHierarchy<T extends WithIdPid>(
    data: T[],
    keyword: string,
    options: {
        fields?: string[]
        exact?: boolean
    } = {}
) {
    const { fields = ['name'], exact = false } = options

    const matched = findMatchedNodes(data, keyword, fields, exact)
    const matchedIds = matched.map((n) => n.id)

    const descendants = findDescendants(data, matchedIds)
    const ancestors = findAncestors(data, matchedIds)

    const all = collectAllLevels(matched, ancestors, descendants)

    return {
        matched, // 匹配的节点
        ancestors, // 所有父级（不包含 matched 自身）
        descendants, // 所有子级
        all // 合并后的完整列表（去重）
    }
}

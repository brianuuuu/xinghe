import { useEffect, useState } from 'react'

type Health = { status: string; service: string; data_dir: string; time: string }

export default function App() {
  const [health, setHealth] = useState<Health | null>(null)
  const [err, setErr] = useState<string>('')

  useEffect(() => {
    fetch('/api/health')
      .then((r) => (r.ok ? r.json() : Promise.reject(new Error(`HTTP ${r.status}`))))
      .then(setHealth)
      .catch((e: Error) => setErr(e.message))
  }, [])

  return (
    <main style={{ fontFamily: 'system-ui, sans-serif', padding: 24, maxWidth: 720 }}>
      <h1 style={{ marginBottom: 8 }}>星河（Xinghe）工作台</h1>
      <p style={{ color: '#6b7280', marginBottom: 16 }}>
        可视化工作台 + 多 harness Agent 协同框架（前端骨架）
      </p>
      <section style={{ padding: 14, border: '1px solid #e5e7eb', borderRadius: 8 }}>
        <strong>后端连接状态（/api/health）：</strong>
        {health ? (
          <pre style={{ marginTop: 8, fontSize: 12 }}>{JSON.stringify(health, null, 2)}</pre>
        ) : (
          <p style={{ marginTop: 8, color: err ? '#dc2626' : '#6b7280' }}>
            {err ? `未连接（${err}），请先运行 go run ./cmd/xinghe` : '检测中…'}
          </p>
        )}
      </section>
    </main>
  )
}

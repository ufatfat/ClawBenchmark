-- 001_create_claw_benchmarks.sql
-- 创建 Claw 测评数据表

CREATE TABLE IF NOT EXISTS claw_benchmarks (
    id               SERIAL PRIMARY KEY,
    name             VARCHAR(100) NOT NULL,
    version          VARCHAR(50)  NOT NULL,
    description      TEXT,

    -- 评分维度 (0-100)
    performance      DECIMAL(5, 2) NOT NULL CHECK (performance >= 0 AND performance <= 100),
    functionality    DECIMAL(5, 2) NOT NULL CHECK (functionality >= 0 AND functionality <= 100),
    usability        DECIMAL(5, 2) NOT NULL CHECK (usability >= 0 AND usability <= 100),
    stability        DECIMAL(5, 2) NOT NULL CHECK (stability >= 0 AND stability <= 100),
    -- 综合评分 = performance*0.3 + functionality*0.3 + usability*0.2 + stability*0.2
    overall_score    DECIMAL(5, 2) NOT NULL CHECK (overall_score >= 0 AND overall_score <= 100),

    -- 性能测试数据
    response_time    INTEGER       NOT NULL CHECK (response_time >= 0),  -- ms
    throughput       INTEGER       NOT NULL CHECK (throughput >= 0),     -- req/s
    memory_usage     INTEGER       NOT NULL CHECK (memory_usage >= 0),   -- MB
    cpu_usage        DECIMAL(5, 2) NOT NULL CHECK (cpu_usage >= 0 AND cpu_usage <= 100),

    -- 功能特性
    features         JSONB,
    supported_skills INTEGER       NOT NULL DEFAULT 0 CHECK (supported_skills >= 0),

    -- 元数据
    test_date        TIMESTAMP     NOT NULL DEFAULT NOW(),
    test_env         VARCHAR(200),
    status           VARCHAR(20)   NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'archived')),

    created_at       TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP     NOT NULL DEFAULT NOW()
);

-- 唯一索引：名称唯一
CREATE UNIQUE INDEX IF NOT EXISTS idx_claw_benchmarks_name ON claw_benchmarks (name);

-- 综合评分索引（用于排序查询）
CREATE INDEX IF NOT EXISTS idx_claw_benchmarks_overall_score ON claw_benchmarks (overall_score DESC);

-- 状态索引
CREATE INDEX IF NOT EXISTS idx_claw_benchmarks_status ON claw_benchmarks (status);

-- 自动更新 updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_claw_benchmarks_updated_at
    BEFORE UPDATE ON claw_benchmarks
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- 插入示例数据
INSERT INTO claw_benchmarks (
    name, version, description,
    performance, functionality, usability, stability,
    overall_score,
    response_time, throughput, memory_usage, cpu_usage,
    features, supported_skills, test_env
) VALUES
(
    'OpenClaw', 'v1.2.0', 'OpenClaw 是一个开源的 Claw 实现，专注于高性能和可扩展性',
    85.0, 90.0, 78.0, 88.0,
    85.0 * 0.3 + 90.0 * 0.3 + 78.0 * 0.2 + 88.0 * 0.2,
    120, 850, 256, 35.5,
    '["tool_use", "streaming", "vision", "function_calling"]', 42,
    'Ubuntu 22.04, 8 Core, 16GB RAM'
),
(
    'ZeroClaw', 'v0.9.5', 'ZeroClaw 轻量级实现，适合资源受限环境',
    72.0, 75.0, 88.0, 80.0,
    72.0 * 0.3 + 75.0 * 0.3 + 88.0 * 0.2 + 80.0 * 0.2,
    95, 1200, 128, 22.0,
    '["tool_use", "streaming"]', 28,
    'Ubuntu 22.04, 4 Core, 8GB RAM'
),
(
    'NanoClaw', 'v2.0.1', 'NanoClaw 企业级 Claw 实现，功能最完整',
    92.0, 95.0, 82.0, 94.0,
    92.0 * 0.3 + 95.0 * 0.3 + 82.0 * 0.2 + 94.0 * 0.2,
    150, 650, 512, 48.0,
    '["tool_use", "streaming", "vision", "function_calling", "computer_use", "batch"]', 68,
    'Ubuntu 22.04, 16 Core, 32GB RAM'
)
ON CONFLICT (name) DO NOTHING;

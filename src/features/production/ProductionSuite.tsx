import { useEffect, useMemo, useState } from "react";
import { useMutation, useQuery } from "@tanstack/react-query";
import {
  Activity,
  Disc3,
  Download,
  FileAudio,
  FileMusic,
  Gauge,
  ListMusic,
  RefreshCw,
  SlidersHorizontal,
  Upload
} from "lucide-react";
import clsx from "clsx";
import { createJob, getJob, getStoredApiBase, Job, setStoredApiBase } from "../../lib/api";
import { readRecentJobs, rememberJob } from "../../lib/recentJobs";

export function ProductionSuite() {
  const [apiBase, setApiBase] = useState(getStoredApiBase());
  const [selectedJobId, setSelectedJobId] = useState<string | null>(readRecentJobs()[0] ?? null);
  const [recentJobs, setRecentJobs] = useState(readRecentJobs());

  const upload = useMutation({
    mutationFn: (file: File) => createJob(file, apiBase),
    onSuccess: (job) => {
      setSelectedJobId(job.id);
      setRecentJobs(rememberJob(job.id));
    }
  });

  const jobQuery = useQuery({
    queryKey: ["job", apiBase, selectedJobId],
    queryFn: () => getJob(selectedJobId!, apiBase),
    enabled: Boolean(selectedJobId),
    refetchInterval: (query) => {
      const status = query.state.data?.status;
      return status === "queued" || status === "running" ? 1500 : false;
    }
  });

  const job = jobQuery.data;
  const apiHost = useMemo(() => {
    try {
      return new URL(apiBase).host;
    } catch {
      return apiBase;
    }
  }, [apiBase]);

  useEffect(() => {
    setStoredApiBase(apiBase);
  }, [apiBase]);

  return (
    <main className="min-h-screen bg-[#f3f1ea] text-ink">
      <header className="border-b border-line bg-panel">
        <div className="mx-auto flex max-w-7xl flex-col gap-5 px-4 py-5 sm:px-6 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <div className="flex items-center gap-3">
              <Disc3 className="h-8 w-8 text-signal" aria-hidden="true" />
              <h1 className="text-2xl font-semibold tracking-normal">Musician Production Suite</h1>
            </div>
            <p className="mt-2 max-w-3xl text-sm text-zinc-600">
              Audio in. BPM, key, chords, stems, cleanup, MIDI, score PDF, and mixdown out.
            </p>
          </div>
          <label className="flex min-w-0 flex-col gap-1 text-xs font-medium uppercase tracking-wide text-zinc-600">
            API endpoint
            <input
              className="focus-ring w-full rounded border border-line bg-white px-3 py-2 text-sm normal-case text-ink shadow-sm lg:w-96"
              value={apiBase}
              onChange={(event) => setApiBase(event.target.value)}
              aria-label="API endpoint"
            />
          </label>
        </div>
      </header>

      <div className="mx-auto grid max-w-7xl gap-5 px-4 py-5 sm:px-6 lg:grid-cols-[360px_minmax(0,1fr)]">
        <section className="space-y-5">
          <UploadPanel
            isUploading={upload.isPending}
            error={upload.error}
            onFile={(file) => upload.mutate(file)}
          />
          <RecentJobs jobs={recentJobs} selected={selectedJobId} onSelect={setSelectedJobId} />
        </section>

        <section className="min-w-0">
          <JobWorkspace
            job={job}
            apiBase={apiBase}
            apiHost={apiHost}
            isLoading={jobQuery.isFetching}
            error={jobQuery.error}
            onRefresh={() => void jobQuery.refetch()}
          />
        </section>
      </div>
    </main>
  );
}

function UploadPanel({
  isUploading,
  error,
  onFile
}: {
  isUploading: boolean;
  error: Error | null;
  onFile: (file: File) => void;
}) {
  return (
    <div className="rounded border border-line bg-white p-4 shadow-sm">
      <div className="flex items-center justify-between">
        <h2 className="text-base font-semibold">New production job</h2>
        <Upload className="h-5 w-5 text-signal" aria-hidden="true" />
      </div>
      <label className="focus-ring mt-4 flex min-h-44 cursor-pointer flex-col items-center justify-center rounded border border-dashed border-zinc-400 bg-panel px-4 text-center hover:border-signal">
        <FileAudio className="mb-3 h-8 w-8 text-zinc-700" aria-hidden="true" />
        <span className="text-sm font-medium">
          {isUploading ? "Uploading..." : "Drop or choose WAV, MP3, FLAC, AIFF, or OGG"}
        </span>
        <span className="mt-1 text-xs text-zinc-600">Max size follows backend configuration</span>
        <input
          className="sr-only"
          type="file"
          accept="audio/*,.wav,.mp3,.flac,.aiff,.ogg"
          disabled={isUploading}
          onChange={(event) => {
            const file = event.target.files?.[0];
            if (file) onFile(file);
            event.currentTarget.value = "";
          }}
        />
      </label>
      {error ? <p className="mt-3 text-sm text-red-700">{error.message}</p> : null}
    </div>
  );
}

function RecentJobs({
  jobs,
  selected,
  onSelect
}: {
  jobs: string[];
  selected: string | null;
  onSelect: (id: string) => void;
}) {
  return (
    <div className="rounded border border-line bg-white p-4 shadow-sm">
      <div className="flex items-center gap-2">
        <ListMusic className="h-5 w-5 text-signal" aria-hidden="true" />
        <h2 className="text-base font-semibold">Recent jobs</h2>
      </div>
      <div className="mt-3 space-y-2">
        {jobs.length === 0 ? (
          <p className="text-sm text-zinc-600">No local job history yet.</p>
        ) : (
          jobs.map((id) => (
            <button
              className={clsx(
                "focus-ring flex w-full items-center justify-between rounded border px-3 py-2 text-left text-sm",
                selected === id ? "border-signal bg-teal-50" : "border-line bg-panel hover:bg-white"
              )}
              key={id}
              type="button"
              onClick={() => onSelect(id)}
            >
              <span className="truncate font-mono">{id}</span>
            </button>
          ))
        )}
      </div>
    </div>
  );
}

function JobWorkspace({
  job,
  apiBase,
  apiHost,
  isLoading,
  error,
  onRefresh
}: {
  job: Job | undefined;
  apiBase: string;
  apiHost: string;
  isLoading: boolean;
  error: Error | null;
  onRefresh: () => void;
}) {
  if (error) {
    return (
      <div className="rounded border border-red-300 bg-white p-5 text-red-800">
        <p className="font-semibold">Backend request failed</p>
        <p className="mt-2 text-sm">{error.message}</p>
      </div>
    );
  }

  if (!job) {
    return (
      <div className="flex min-h-[520px] items-center justify-center rounded border border-line bg-white p-8 text-center shadow-sm">
        <div>
          <SlidersHorizontal className="mx-auto h-10 w-10 text-signal" aria-hidden="true" />
          <h2 className="mt-4 text-xl font-semibold">Ready for an audio file</h2>
          <p className="mt-2 max-w-xl text-sm text-zinc-600">
            Connected target: {apiHost}. Upload a track to start the production pipeline.
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-5">
      <div className="rounded border border-line bg-white p-5 shadow-sm">
        <div className="flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
          <div className="min-w-0">
            <p className="text-xs font-medium uppercase tracking-wide text-zinc-600">{job.id}</p>
            <h2 className="mt-1 truncate text-xl font-semibold">{job.filename}</h2>
            <p className="mt-2 text-sm text-zinc-600">{job.message}</p>
          </div>
          <button
            className="focus-ring inline-flex items-center justify-center gap-2 rounded border border-line bg-panel px-3 py-2 text-sm font-medium hover:bg-white"
            type="button"
            onClick={onRefresh}
          >
            <RefreshCw className={clsx("h-4 w-4", isLoading && "animate-spin")} aria-hidden="true" />
            Refresh
          </button>
        </div>
        <div className="mt-5 h-3 overflow-hidden rounded bg-zinc-200" aria-label="Job progress">
          <div
            className="h-full bg-signal transition-all"
            style={{ width: `${job.progress}%` }}
          />
        </div>
        <div className="mt-3 flex flex-wrap items-center gap-3 text-sm">
          <StatusPill status={job.status} />
          <span>{job.progress}%</span>
          {job.error ? <span className="text-red-700">{job.error}</span> : null}
        </div>
      </div>

      <div className="grid gap-5 lg:grid-cols-3">
        <Metric icon={Gauge} label="BPM" value={job.bpm?.toFixed(1) ?? "Pending"} />
        <Metric icon={FileMusic} label="Key" value={job.key ?? "Pending"} />
        <Metric icon={Activity} label="Artifacts" value={String(job.artifacts.length)} />
      </div>

      <div className="grid gap-5 lg:grid-cols-[1fr_1.25fr]">
        <div className="rounded border border-line bg-white p-4 shadow-sm">
          <h3 className="text-base font-semibold">Chord timeline</h3>
          <div className="mt-3 max-h-80 overflow-auto">
            {job.chords.length === 0 ? (
              <p className="text-sm text-zinc-600">No chord estimates yet.</p>
            ) : (
              <table className="w-full text-sm">
                <thead className="text-left text-xs uppercase tracking-wide text-zinc-600">
                  <tr>
                    <th className="py-2">Start</th>
                    <th className="py-2">End</th>
                    <th className="py-2">Chord</th>
                  </tr>
                </thead>
                <tbody>
                  {job.chords.map((chord) => (
                    <tr className="border-t border-line" key={`${chord.start}-${chord.chord}`}>
                      <td className="py-2">{chord.start.toFixed(1)}s</td>
                      <td className="py-2">{chord.end.toFixed(1)}s</td>
                      <td className="py-2 font-semibold">{chord.chord}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            )}
          </div>
        </div>

        <div className="rounded border border-line bg-white p-4 shadow-sm">
          <h3 className="text-base font-semibold">Downloads</h3>
          <div className="mt-3 grid gap-2">
            {job.artifacts.length === 0 ? (
              <p className="text-sm text-zinc-600">Artifacts appear as processing finishes.</p>
            ) : (
              job.artifacts.map((artifact) => (
                <a
                  className="focus-ring flex items-center justify-between gap-3 rounded border border-line bg-panel px-3 py-2 text-sm hover:bg-white"
                  href={`${apiBase}${artifact.url}`}
                  key={artifact.url}
                >
                  <span className="min-w-0">
                    <span className="block truncate font-medium">{artifact.name}</span>
                    <span className="text-xs text-zinc-600">
                      {artifact.kind} · {formatBytes(artifact.sizeBytes)}
                    </span>
                  </span>
                  <Download className="h-4 w-4 shrink-0 text-signal" aria-hidden="true" />
                </a>
              ))
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

function StatusPill({ status }: { status: Job["status"] }) {
  const classes = {
    queued: "border-zinc-300 bg-zinc-100 text-zinc-800",
    running: "border-teal-300 bg-teal-50 text-teal-800",
    succeeded: "border-green-300 bg-green-50 text-green-800",
    failed: "border-red-300 bg-red-50 text-red-800"
  }[status];
  return <span className={clsx("rounded border px-2 py-1 text-xs font-semibold", classes)}>{status}</span>;
}

function Metric({
  icon: Icon,
  label,
  value
}: {
  icon: typeof Gauge;
  label: string;
  value: string;
}) {
  return (
    <div className="rounded border border-line bg-white p-4 shadow-sm">
      <div className="flex items-center gap-2 text-sm text-zinc-600">
        <Icon className="h-4 w-4 text-signal" aria-hidden="true" />
        {label}
      </div>
      <p className="mt-2 text-2xl font-semibold">{value}</p>
    </div>
  );
}

function formatBytes(bytes: number): string {
  if (bytes < 1024) return `${bytes} B`;
  if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(1)} KB`;
  return `${(bytes / (1024 * 1024)).toFixed(1)} MB`;
}

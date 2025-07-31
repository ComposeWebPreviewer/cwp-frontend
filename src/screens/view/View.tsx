import { useParams } from "react-router"
import type { CodespaceDetailsResponse } from "../../response/CodespaceDetailsResponse"
import { useQuery } from "@tanstack/react-query"
import { API_BASE_URL } from "../../App"
import type { BaseResponse } from "../../response/BaseResponse"

export default function View() {
  const params = useParams<{ id: string }>()
  const { isPending, error, data } = useQuery({
    queryKey: [params.id],
    queryFn: async () => {
      const res = await fetch(`${API_BASE_URL}/?id=${params.id}`)
      const body: BaseResponse = await res.json()

      if (body.status == 200) {
        const details: CodespaceDetailsResponse = await res.json()
        return details
      }
    }
  })

  if (isPending || !data || error) {
    return <div>Loading...</div>
  }

  const handleRunClick = () => {
    loadPreview(data.wasm);
  };

  return (
    <>
      <script src={data.js}></script>
      <div className="split-screen-container">
        <div className="header">
          <button className="run-button" onClick={handleRunClick}>
            Run
          </button>
        </div>
        <div className="panes-container">
          <div className="pane left-pane">
            {/* Content for the left pane */}
            <h2>Left Pane</h2>
            <p>This is the content for the left side.</p>
          </div>
          <div className="pane right-pane composeApp">
            {/* Content for the right pane */}
            <h2>Right Pane</h2>
            <p>This is the content for the right side.</p>
          </div>
        </div>
      </div>
      <div>
        <h1>View Page</h1>
        <p>This is the view page content.</p>
      </div>
    </>
  )
}

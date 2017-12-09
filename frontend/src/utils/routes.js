var routes = {
  "Landing": "/",
  "Classes": "/classes",
  "Login as Student": "/login",
  "Login as Professor": "/login",
  "Login": "/login",
  "Projects": "/classes/:class_id",
  "Analytics": "/classes/:class_id/projects/:project_id",
  "Submission": "/classes/:class_id/projects/:project_id",
  "Create Class": "/classes/create",
  "Edit Class": "/classes/:class_id/edit",
  "Create Project": "/classes/:class_id/projects/create",
  "Edit Project": "/classes/:class_id/projects/:project_id/edit",
}

export default routes
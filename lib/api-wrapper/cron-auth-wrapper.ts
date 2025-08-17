export function withCronAuth(handler: (req: Request) => Promise<Response>) {
  return async (req: Request): Promise<Response> => {
    const authHeader = req.headers.get("authorization");

    if (authHeader !== `Bearer ${process.env.CRON_TOKEN}`) {
      return Response.json({
        message: "unauthorized",
      }, {
        status: 401,
      });
    }

    return handler(req);
  };
}

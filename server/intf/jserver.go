package intf

type XServer interface {
	/**

	 */
	Start()

	Serve()

	Stop()

	AddRouter(r *XRouter)
}

import utils
import deployment_options
import pvc_size_utils


log = utils.get_logger('deploy_postgres')


def main():
    deploy_options = deployment_options.load_deployment_options()

    log.info('Starting postgres deployment')

    utils.verify_build_directory(deploy_options.namespace)

    deploy_postgres_secret(deploy_options)

    log.info('Completed postgres deployment')


def deploy_postgres_secret(deploy_options):
    docs = utils.load_yaml_file_docs('deploy/postgres/postgres-secret.yaml')

    utils.set_namespace_in_yaml_docs(docs, deploy_options.namespace)

    dst_file = utils.dump_yaml_file_docs(
        basename=f'build/{deploy_options.namespace}/postgres-secret.yaml',
        docs=docs
    )

    log.info('Deploying %s', dst_file)
    utils.apply(
        target=deploy_options.target,
        namespace=deploy_options.namespace,
        profile=deploy_options.profile,
        file=dst_file
    )


if __name__ == "__main__":
    main()

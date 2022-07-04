<template>
    <div>
        <div class="offcanvas offcanvas-start" data-bs-scroll="true" data-bs-backdrop="false" tabindex="-1"
            id="offcanvasScrolling" aria-labelledby="offcanvasScrollingLabel">
            <nav class="navbar-dark bg-image text-white h-100">

                <div class="offcanvas-header">
                    <h4 class="offcanvas-title">
                        <a th:href="#" class="text-white">
                            <span class="d-inline-block me-1" style="transform: rotate(-15deg)">
                                <i class="bi bi-credit-card-2-front-fill fs-3">
                                </i>
                            </span>
                            ユーザ管理
                        </a>
                    </h4>
                </div>

                <div class="offcanvas-body pt-0">
                    <hr class="mt-1">
                    <ul class="navbar-nav">
                        <li class="nav-item">
                            <a th:href="#" class="nav-link text-white">
                                <i class="bi bi-person-fill me-2"></i>
                                ログイン情報
                            </a>
                        </li>
                        <li class="nav-item">
                            <a th:href="#" class="nav-link text-white">
                                <i class="bi bi-box-arrow-right me-2"></i>
                                ログアウト
                            </a>
                        </li>
                    </ul>

                    <hr>
                    <div class="text-center d-lg-none">
                        <button type="button" class="btn-circle btn-light" data-bs-dismiss="offcanvas"
                            aria-label="Close">
                            <i class="bi bi-chevron-left"></i>
                        </button>
                    </div>
                </div>
            </nav>
        </div>
    </div>
    <div class="toggler">
        <nav class="navbar navbar-expand-lg navbar-light">
            <div class="container-fluid">
                <button class="navbar-toggler" type="button" data-bs-toggle="offcanvas"
                    data-bs-target="#offcanvasScrolling" aria-controls="offcanvasScrolling" aria-expanded="false"
                    aria-label="Toggle offcanvas">
                    <span class="navbar-toggler-icon"></span>
                </button>
            </div>
        </nav>
    </div>
</template>
<script>
export default {

}
</script>
<style scoped>
.navbar-dark {
    background-color: blueviolet;
}

.offcanvas {
    width: 200px;
}

.btn-circle {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    opacity: 0.5;
    border: 0;
}
</style>

<style>
@media (min-width: 992px) {
    .offcanvas {
        transform: none;
        visibility: visible !important;
    }

    .offcanvas-outside {
        transition: margin-left 0.3s ease-in-out;
        margin-left: 200px;
    }
}
</style>